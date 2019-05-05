clear all;
close all;
clc;

ensembles = 10000;
betaLowerLim = 20;
betaUpperLim = 70;
betaStep = 10;

figure;
hold on;

for beta = betaLowerLim:betaStep:betaUpperLim
    Sweeps = [];
    % Betas = [];
    % SubSystemIDs = [];
    Magnetisations = [];
    DimensionlessEnergies = [];

    i = 0;
    while( i < ensembles )
        fname = ['./section1final/beta-', num2str(beta) ,'-system', num2str(i) ,'.csv'];
        data = csvread(fname, 1);

        Sweep = data(:,1);
    %     Beta = data(:,2);
    %     SubSystemID = data(:,3);
%         Magnetisation = data(:,4);
        DimensionlessEnergy = data(:,5);

        Sweeps = [Sweeps Sweep];
    %     Betas = [Betas Beta];
    %     SubSystemIDs = [SubSystemIDs SubSystemID];
%         Magnetisations = [Magnetisations Magnetisation];
        DimensionlessEnergies = [DimensionlessEnergies DimensionlessEnergy];

        i = i+1;
    end

    meanDimensionlessEnergy = mean(DimensionlessEnergies, 2)
    standDevInDimensionlessEnergy = std(DimensionlessEnergies, 0, 2);
%     plot(Sweeps(:,1), meanMagnetisations)
%     errorbar(Sweeps(:,1), meanMagnetisations, standDevInMagnetisations, 'LineStyle', 'none')
    errorbar(Sweeps(:,1), meanDimensionlessEnergy, standDevInDimensionlessEnergy, 'x-', 'DisplayName', sprintf('$\\beta = %0.2f$', beta/100))

end

legend_handle = legend('-DynamicLegend');
legend_handle.FontSize = 22;
legend_handle.Orientation = 'horizontal';
legend_handle.Location = 'northoutside';
set(legend_handle,'Interpreter','latex')
legend('show');

ax = gca();
ax.FontSize = 15;

hold off;
xlabel('Number of sweeps', 'FontSize', 16);
ylabel('$E/J$', 'Interpreter', 'latex', 'FontSize', 16);
% title('Ensemble average of n=10000 independent clusters')